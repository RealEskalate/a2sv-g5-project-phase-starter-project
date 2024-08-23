

import 'package:equatable/equatable.dart';
abstract class ProductEvent  extends Equatable{
  const ProductEvent ();

  @override
  List<Object ?> get props => [];
}

class LoadAllProductEvent extends ProductEvent {
  const LoadAllProductEvent ();

  @override
  List<Object ?> get props => [];
}

class GetSingleProductEvent extends ProductEvent {
  final String id;
  const GetSingleProductEvent({
    required this.id
  });
  @override
  List<Object ?> get props => [id];
}


class DeleteProductEvent extends ProductEvent {
  final String id;
  const DeleteProductEvent({
    required this.id
  });
  @override
  List<Object ?> get props => [id];
}


class CreateProductEvent extends ProductEvent {
  final String ecommerceEntity;
  const CreateProductEvent({
    required this.ecommerceEntity
  });
  @override
  List<Object ?> get props => [ecommerceEntity];
}








