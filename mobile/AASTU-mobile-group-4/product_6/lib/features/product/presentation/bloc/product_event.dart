// product_events.dart

import 'package:equatable/equatable.dart';
import 'package:meta/meta.dart';

import '../../domain/entities/product.dart';

@immutable
abstract class ProductEvent extends Equatable {
  const ProductEvent();

  @override
  List<Object> get props => [];
}

// Event to load all products
class LoadAllProductEvent extends ProductEvent {}

// Event to retrieve a single product by its ID
class GetSingleProductEvent extends ProductEvent {
  final String productId;

  const GetSingleProductEvent(this.productId);

  @override
  List<Object> get props => [productId];
}

// Event to update a product's details
class UpdateProductEvent extends ProductEvent {
  final Product product;

  const UpdateProductEvent(this.product);

  @override
  List<Object> get props => [product];
}

// Event to delete a product
class DeleteProductEvent extends ProductEvent {
  final String productId;

  const DeleteProductEvent(this.productId);

  @override
  List<Object> get props => [productId];
}

// Event to create a new product
class CreateProductEvent extends ProductEvent {
  final Product product;
  

  const CreateProductEvent({required this.product});

  @override
  List<Object> get props => [product];
}
