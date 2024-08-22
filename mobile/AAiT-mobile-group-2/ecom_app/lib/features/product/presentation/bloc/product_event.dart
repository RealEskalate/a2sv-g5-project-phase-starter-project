part of 'product_bloc.dart';

abstract class ProductEvent extends Equatable {
  const ProductEvent();

  @override
  List<Object> get props => [];
}

class LoadAllProductEvent extends ProductEvent {}

class GetSingleProductEvent extends ProductEvent {
  final String id;

  GetSingleProductEvent({required this.id});


  @override
  List<Object> get props => [id];
}

class UpdateProductEvent extends ProductEvent {
  final Product product;

  UpdateProductEvent({required this.product});


  @override
  List<Object> get props => [product];
}

class DeleteProductEvent extends ProductEvent {
  final String id;

  DeleteProductEvent({required this.id});


  @override
  List<Object> get props => [id];
}


class CreateProductEvent extends ProductEvent {
  final Product product;

  CreateProductEvent({required this.product});


  @override
  List<Object> get props => [product];
}
