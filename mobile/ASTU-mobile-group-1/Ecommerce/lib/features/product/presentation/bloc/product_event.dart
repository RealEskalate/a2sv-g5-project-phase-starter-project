part of 'product_bloc.dart';

sealed class ProductEvent {
  const ProductEvent();
}

class LoadAllProductEvent extends ProductEvent {
  LoadAllProductEvent();
}

class GetSingleProductEvent extends ProductEvent {
  final String id;
  GetSingleProductEvent({required this.id});
}

class UpdateProductEvent extends ProductEvent {
  final String id;
  final String name;
  final String description;
  final String price;
  final String imageUrl;

  UpdateProductEvent(
      {required this.id,
      required this.name,
      required this.description,
      required this.price,
      required this.imageUrl});
}

class DeleteProductEvent extends ProductEvent {
  final String id;

  DeleteProductEvent({required this.id});
}

class CreateProductEvent extends ProductEvent {
  final String id;
  final String name;
  final String description;
  final String price;
  final String imageUrl;

  CreateProductEvent(
      {required this.id,
      required this.name,
      required this.description,
      required this.price,
      required this.imageUrl});
}

class UpdateTextFieldEvent extends ProductEvent {
  final String name;
  final String value;

  const UpdateTextFieldEvent({required this.name, required this.value});
}

class AddProductEvent extends ProductEvent {}

class ResetMessageStateEvent extends ProductEvent {}

class FilterProductEvent extends ProductEvent {
  final String text;

  const FilterProductEvent({required this.text});
}
