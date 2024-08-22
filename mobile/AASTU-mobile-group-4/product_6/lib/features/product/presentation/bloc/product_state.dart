// product_states.dart

import 'package:equatable/equatable.dart';
import 'package:meta/meta.dart';

import '../../domain/entities/product.dart';

@immutable
abstract class ProductState extends Equatable {
  const ProductState();

  @override
  List<Object> get props => [];
}

// Initial state before any data is loaded
class InitialState extends ProductState {}

// State when the app is currently fetching data
class LoadingState extends ProductState {}

// State when all products are successfully loaded
class LoadedAllProductState extends ProductState {
  final List<Product> products;

  const LoadedAllProductState(this.products);

  @override
  List<Object> get props => [products];
}

// State when a single product is successfully retrieved
class LoadedSingleProductState extends ProductState {
  final Product product;

  const LoadedSingleProductState({required this.product});

  @override
  List<Object> get props => [product];
}

// State when an error has occurred during data retrieval or processing
class ErrorState extends ProductState {
  final String message;

  const ErrorState(this.message);

  @override
  List<Object> get props => [message];
}

class Success extends ProductState {
  final String message;

  const Success({required this.message });

  @override
  List<Object> get props => [message];
}


//State 

// class EditState extends ProductState{
//   final
// }