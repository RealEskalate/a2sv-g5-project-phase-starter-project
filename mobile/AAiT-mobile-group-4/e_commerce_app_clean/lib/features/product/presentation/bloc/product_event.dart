part of 'product_bloc.dart';
// LoadAllProductEvent: This event should be dispatched when the user wants to load all products from the repository.
// LoadSingleProductEvent: Dispatch this event when the user wants to retrieve a single product using its ID.
// UpdateProductEvent: Dispatch this event when the user wants to update a product's details.
// DeleteProductEvent: Dispatch this event when the user wants to delete a product.
// CreateProductEvent: Dispatch this event when the user wants to create a new product.

sealed class ProductEvent extends Equatable {
  const ProductEvent();

  @override
  List<Object> get props => [];
}

class LoadAllProductEvent extends ProductEvent {
  @override
  List<Object> get props => [];
}

class LoadSingleProductEvent extends ProductEvent {
  final String id;

  const LoadSingleProductEvent({required this.id});

  @override
  List<Object> get props => [id];
}

class UpdateProductEvent extends ProductEvent {
  final ProductEntity product;

  const UpdateProductEvent({required this.product});

  @override
  List<Object> get props => [product];
}

class DeleteProductEvent extends ProductEvent {
  final String id;

  const DeleteProductEvent({required this.id});

  @override
  List<Object> get props => [id];
}

class CreateProductEvent extends ProductEvent {
  final ProductEntity product;

  const CreateProductEvent({required this.product});

  @override
  List<Object> get props => [product];
}
