module Services
  module Api 
    module V1
      module Products
        class Upsert
          attr_accessor :params, :request

          def initialize(params, request)
            @params = params
            @request = request
          end

          def execute
            raise ActiveRecord::RecordNotFound, "Product Not Found" if product.blank?

            ActiveRecord::Base.transaction do
              product.id        ||= params[:id]           if params[:id].present?
              product.name        = params[:name]         if params[:name].present?
              product.brand       = params[:brand]        if params[:brand].present?
              product.price       = params[:price]        if params[:price].present?
              product.description = params[:description]  if params[:description].present?
              product.stock       = params[:stock]        if params[:stock].present?

              product.save!

              if !!params[:is_api]
                begin
                  Timeout.timeout(5) do
                    Karafka.producer.produce_sync(topic: 'rails-to-go', payload: product.to_json)
                  end
                rescue Timeout::Error
                  Rails.logger.error("Failed to deliver message to Kafka: Timeout error")
                rescue Kafka::DeliveryFailed => e
                  Rails.logger.error("Failed to deliver message to Kafka: #{e.message}")
                rescue StandardError => e
                  Rails.logger.error("Failed to deliver message to Kafka: #{e.message}")
                end
              end
            end
            
            product
          end

          private
          def product
            @product ||= Product.find_by(id: params[:id]).presence || Product.new
          end
        end
      end
    end
  end
end