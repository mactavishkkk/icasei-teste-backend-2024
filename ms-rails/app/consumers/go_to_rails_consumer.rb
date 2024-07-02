class GoToRailsConsumer < ApplicationConsumer
    def consume
      messages.each do |message|
        product_params = message.payload.with_indifferent_access
  
        product = Product.find_by(id: product_params[:id]) || Product.new
        product.assign_attributes(product_params.except(:created_at, :updated_at))
  
        if product.save
          Rails.logger.info("Product created/updated successfully: #{product.inspect}")
        else
          Rails.logger.error("Failed to create/update product: #{product.errors.full_messages}")
        end
      end
    end
  end
  