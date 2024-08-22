part of 'update_page_bloc.dart';

abstract class UpdatePageEvent extends Equatable {
  const UpdatePageEvent();

  @override
  List<Object> get props => [];
}

class UpdateProductEvent extends UpdatePageEvent {
  final ProductModel product;

  const UpdateProductEvent(this.product);

  @override
  List<Object> get props => [product];
}

class DeleteProductEvent extends UpdatePageEvent {
  final String productId;

  const DeleteProductEvent(this.productId);

  @override
  List<Object> get props => [productId];
}
