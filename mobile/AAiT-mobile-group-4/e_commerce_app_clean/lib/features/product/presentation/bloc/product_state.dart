part of 'product_bloc.dart';

sealed class ProductState extends Equatable {
  const ProductState();

  @override
  List<Object> get props => [];
}

final class ProductInitial extends ProductState {}

final class ProductLoading extends ProductState {}

final class LoadedSingleProductState extends ProductState {
  final ProductEntity product;

  const LoadedSingleProductState(this.product);

  @override
  List<Object> get props => [product];
}

class LoadedAllProductState extends ProductState {
  final List<ProductEntity> products;

  const LoadedAllProductState(this.products);

  @override
  List<Object> get props => [products];
}

class ProductErrorState extends ProductState {
  final String message;

  const ProductErrorState(this.message);

  @override
  List<Object> get props => [message];
}

class ProductDeletedState extends ProductState {
  final String message;
  const ProductDeletedState({required this.message});

  @override
  List<Object> get props => [message];
}

class ProductUpdatedState extends ProductState {
  final ProductEntity product;
  const ProductUpdatedState(this.product);

  @override
  List<Object> get props => [product];
}

class ProductCreatedState extends ProductState {
  final ProductEntity product;
  const ProductCreatedState(this.product);

  @override
  List<Object> get props => [product];
}
