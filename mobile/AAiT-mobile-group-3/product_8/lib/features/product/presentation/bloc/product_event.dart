

import 'package:equatable/equatable.dart';

import '../../domain/entities/product_entity.dart';
abstract class ProductEvent extends Equatable{
  const ProductEvent();

  @override
  List<Object?> get props => [];
}


class LoadProduct extends ProductEvent {}


class CreateProductEvent extends ProductEvent{
  final Product product;
  const CreateProductEvent({required this.product});

  @override 
  List<Object?> get props => [product];
}

class UpdateProductEvent extends ProductEvent{
  final Product product;
  const UpdateProductEvent({required this.product});

  @override 
  List<Object?> get props => [product];
}

class DeleteProductEvent extends ProductEvent{
  final String id;
  const DeleteProductEvent({required this.id});

  @override 
  List<Object?> get props => [id];
}

class GetProductByIdEvent extends ProductEvent{
  final String id;
  const GetProductByIdEvent({required this.id});

  @override 
  List<Object?> get props => [id];
}

