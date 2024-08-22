part of 'add_page_bloc.dart';

abstract class AddPageEvent extends Equatable {
  const AddPageEvent();

  @override
  List<Object> get props => [];
}

class AddProductEvent extends AddPageEvent {
  final ProductModel product;
  final String imagePath;

  const AddProductEvent(this.product, this.imagePath);

  @override
  List<Object> get props => [product, imagePath];
}

class UpdateProductEvent extends AddPageEvent {
  final ProductModel product;

  const UpdateProductEvent(this.product);

  @override
  List<Object> get props => [product];
}
