import 'package:equatable/equatable.dart';

import '../../domain/entitity/product.dart';

abstract class ProductState extends Equatable {
  const ProductState();

  @override
  List<Object> get props => [];
}

class ProductInitial extends ProductState {}

class ProductStateEmpty extends ProductState {}

class ProductStateLoading extends ProductState {}

class ProductStateLoaded extends ProductState {
  final Product product;
  const ProductStateLoaded({required this.product});
  @override
  List<Object> get props => [product];
}

class ProductLoadFailure extends ProductState {
  final String message;
  const ProductLoadFailure({required this.message});
  @override
  List<Object> get props => [message];
}

class AllProductsLoaded extends ProductState {
  final List<Product> products;
  const AllProductsLoaded({required this.products});
  @override
  List<Object> get props => [products];
}

class AllProductsLoadedFailure extends ProductState {
  final String message;
  const AllProductsLoadedFailure({required this.message});
  @override
  List<Object> get props => [message];
}

class ProductDeleteState extends ProductState {
  final Product product;
  const ProductDeleteState({required this.product});
  @override
  List<Object> get props => [product];
}


class ProductDeleteFailureState extends ProductState {
  final String message;
  const ProductDeleteFailureState({required this.message});
  @override
  List<Object> get props => [message];
}

class ProductUpdateState extends ProductState {
  final Product product;
  const ProductUpdateState({required this.product});
  @override
  List<Object> get props => [product];
}


class ProductUpdateFailureState extends ProductState {
  final String message;
  const ProductUpdateFailureState({required this.message});
  @override
  List<Object> get props => [message];
}

class ProductInsertState extends ProductState {
  final Product product;
  const ProductInsertState({required this.product});
  @override
  List<Object> get props => [product];
}

class ProductInsertFailureState extends ProductState {
  final String message;
  const ProductInsertFailureState({required this.message});
  @override
  List<Object> get props => [message];
}