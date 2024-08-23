


import 'package:equatable/equatable.dart';

import '../../../Domain/entity/ecommerce_entity.dart';

abstract class ProductState extends Equatable{
  const ProductState();

  @override
  List<Object ?> get props => [];
}

class ProductIntialState extends ProductState {}

class LoadedSingleProductState extends ProductState {
  final EcommerceEntity product;

  const LoadedSingleProductState ({
    required this.product
  });

  @override
  List<Object ?> get props => [product];
}

class LoadedAllProductState extends ProductState {
  final List<EcommerceEntity> products;

  const LoadedAllProductState ({
    required this.products
  });

  @override

  List<Object ?> get props => [products]; 
}

class SuccessDelete extends ProductState {
  final bool deleted;

  const SuccessDelete ({
    required this.deleted
  });

  @override

  List<Object?> get props => [deleted];
}


class SuccessAdd extends ProductState {
  final bool add;

  const SuccessAdd ({
    required this.add
  });

  @override

  List<Object?> get props => [add];
}



class ProductErrorState extends ProductState {
  final String messages;

  const ProductErrorState ({
    required this.messages
  });

  @override
  List<Object ?> get props => [messages];
}




class LoadingState extends ProductState {}

