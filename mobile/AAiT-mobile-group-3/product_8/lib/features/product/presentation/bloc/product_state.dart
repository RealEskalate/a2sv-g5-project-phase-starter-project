


import 'package:equatable/equatable.dart';

import '../../domain/entities/product_entity.dart';



abstract class ProductState extends Equatable {
  const ProductState();

  @override
  List<Object> get props => [];
}

 class ProductInitial extends ProductState {}

 class ProductLoading extends ProductState {}

  class ProductLoaded extends ProductState {
    final List<Product> products;
  
    const ProductLoaded({required this.products});
  
    @override
    List<Object> get props => [products];
  }


  class ProductError extends ProductState {
    
  }


class ProductLoadedSingle extends ProductState {
  final Product product;

  const ProductLoadedSingle({required this.product});

  @override
  List<Object> get props => [product];
}






class ProductDeleteState extends ProductState {
  
}

class ProductUpdatedState extends ProductState {
  final Product product;

  const ProductUpdatedState({required this.product});

  @override
  List<Object> get props => [product];
}

class ProductCreatedState extends ProductState {
  final Product product;

  const ProductCreatedState({required this.product});

  @override
  List<Object> get props => [product];
}

class ProductUpdatedErrorState extends ProductState {
  final String message;

  const ProductUpdatedErrorState({required this.message});

  @override
  List<Object> get props => [message];
}

class ProductCreatedErrorState extends ProductState {
  final String message;

  const ProductCreatedErrorState({required this.message});

  @override
  List<Object> get props => [message];
}

