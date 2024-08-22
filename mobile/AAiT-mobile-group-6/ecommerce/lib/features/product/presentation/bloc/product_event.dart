import 'package:equatable/equatable.dart';

import '../../domain/entitity/product.dart';

abstract class ProductEvent extends Equatable {
  const ProductEvent();

  @override
  List<Object> get props => [];
}

class GetProductEvent extends ProductEvent {
  final String productId;
  const GetProductEvent({required this.productId});
  @override
  List<Object> get props => [productId];
}

class GetAllProductEvent extends ProductEvent {}

class InsertProductEvent extends ProductEvent {
  final Product product;

  InsertProductEvent({required this.product});

  @override
  List<Object> get props => [product];
}

class UpdateProductEvent extends ProductEvent {
  final Product product;
  UpdateProductEvent({required this.product});

  @override
  List<Object> get props => [product];
}

class DeleteProductEvent extends ProductEvent {
  final String productId;
  DeleteProductEvent({required this.productId});

  @override
  List<Object> get props => [productId];
}
