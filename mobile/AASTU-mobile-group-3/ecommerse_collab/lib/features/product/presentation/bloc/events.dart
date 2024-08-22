import '../../domain/entity/product.dart';

class ProductEvent {
  ProductEvent();
}

// LoadAllProductEvent
class GetAllProductEvent extends ProductEvent{}

// GetSingleProductEvent
class GetProductEvent extends ProductEvent{
  final String productId;

  GetProductEvent({required this.productId});
}

// UpdateProductEvent
class UpdateProductEvent  extends ProductEvent{
  final String productId;
  final String newName;
  final int newPrice;
  final String newDescription;

  UpdateProductEvent({required this.productId, required this.newName,   required this.newPrice, required this.newDescription});
}

// DeleteProductEvent
class DeleteProductEvent  extends ProductEvent{
  final String productId;

  DeleteProductEvent({required this.productId});
}

// CreateProductEvent
class AddProductEvent  extends ProductEvent{
  final Product product;

  AddProductEvent({required this.product});
}
