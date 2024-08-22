part of 'product_bloc.dart';

abstract class ProductState extends Equatable {
  const ProductState();

  @override
  List<Object> get props => [];
}

class ProductInitialState extends ProductState {}

class ProductLoading extends ProductState {}

class LoadAllProductState extends ProductState {
  final List<Product> products;

  LoadAllProductState({required this.products});

  @override
  List<Object> get props => [products];
}

class LoadSingleProductState extends ProductState {
  final Product product;

  LoadSingleProductState({required this.product});

  @override
  List<Object> get props => [product];
}

class ProductErrorState extends ProductState {
  final String message;

  ProductErrorState({required this.message});

  @override
  List<Object> get props => [message];
}

class ProductDeletedState extends ProductState {}

class ProductUpdatedState extends ProductState {
  final Product product;

  ProductUpdatedState({required this.product});

  @override
  List<Object> get props => [product];
}

class ProductCreatedState extends ProductState {
  final Product product;

  ProductCreatedState({required this.product});

  @override
  List<Object> get props => [product];
}

class ProductUpdatedErrorState extends ProductState {
  final String message;

  ProductUpdatedErrorState({required this.message});

  @override
  List<Object> get props => [message];
}

class ProductCreatedErrorState extends ProductState {
  final String message;

  ProductCreatedErrorState({required this.message});

  @override
  List<Object> get props => [message];
}
