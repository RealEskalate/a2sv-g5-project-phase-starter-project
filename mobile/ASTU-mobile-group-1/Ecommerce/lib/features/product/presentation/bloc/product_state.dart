part of 'product_bloc.dart';

sealed class ProductState extends Equatable {
  const ProductState();

  @override
  List<Object> get props => [];
}

class InitalState extends ProductState {}

class LoadingState extends ProductState {}

class LoadedAllProductsState extends ProductState {
  final List<ProductEntity> products;

  const LoadedAllProductsState({required this.products});

  @override
  List<Object> get props => products;
}

class LoadedSingleProductState extends ProductState {
  final ProductEntity product;

  const LoadedSingleProductState({required this.product});
  @override
  List<Object> get props => [product];
}

class ErrorState extends ProductState {
  final String message;

  const ErrorState({required this.message});
  @override
  List<Object> get props => [message];
}

class UpdatedProductState extends ProductState {}

class DeletedProductState extends ProductState {}

class DeletingProductState extends ProductState {}

class AddedProductState extends ProductState {}

class AddProuctState extends ProductState {}

class NeutralState extends ProductState {}

class ShowMessageState extends ProductState {
  final String message;
  const ShowMessageState({required this.message});
}
