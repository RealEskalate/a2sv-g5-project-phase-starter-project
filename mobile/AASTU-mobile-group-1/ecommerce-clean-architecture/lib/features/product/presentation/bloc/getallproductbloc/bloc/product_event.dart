

part of 'product_bloc.dart';

abstract class ProductEvent extends Equatable {
  const ProductEvent();
  @override
  List<Object> get props => [];
}
class GetAllProductEvent extends ProductEvent{
  @override
  List<Object> get props => [];
}

class AddProductEvent extends ProductEvent {
  final ProductModel product;
  const AddProductEvent({required this.product});
  @override
  List<Object> get props => [product];
} 

class UpdateProductEvent extends ProductEvent {
  final ProductModel product;
  UpdateProductEvent(
    {required this.product});
  @override
  List<Object> get props => [product];
}

class DeleteProductEvent extends ProductEvent {
  final String id;
  DeleteProductEvent(this.id);
  @override
  List<Object> get props => [id];
}

class GetProductEvent extends ProductEvent{
  final String id;
  GetProductEvent(this.id);
  @override
  List<Object> get props => [id];
}