// Define the abstract base class for all states

import '../../domain/entity/product.dart';

abstract class ProductState {}

// Represents the initial state before any data is loaded
class InitialState extends ProductState {}

// Indicates that the Product is currently fetching data
class LoadingState extends ProductState {}
class AddingState extends ProductState {}
class LoadedState extends ProductState {
  final List<Product> products;

  LoadedState(this.products);
}

// Represents the state where all products are successfully loaded from the repository
class GetAllProductState extends ProductState {
  final List<Product> products;

  GetAllProductState(this.products);
}

// Represents the state where a single product is successfully retrieved
class GetProductState extends ProductState {
  final Product product;

  GetProductState(this.product);
}

// Indicates that an error has occurred during data retrieval or processing
class ErrorState extends ProductState {
  final String message;

  ErrorState(this.message);
}

class SuccessState extends ProductState {
  final String message;

  SuccessState(this.message);
}

class AddProductState extends ProductState {
  final Product product;

  AddProductState(this.product);
}

class UpdateProductState extends ProductState {
  final Product product;

  UpdateProductState(this.product);
}

