

import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';

abstract class ProductState {}



class ProductLoading extends ProductState {}

class ProductAddedSuccess extends ProductState {
  final String message;

  ProductAddedSuccess({required this.message});
}

class ProductAddedFailure extends ProductState {
  final String error;

  ProductAddedFailure({required this.error});
}
