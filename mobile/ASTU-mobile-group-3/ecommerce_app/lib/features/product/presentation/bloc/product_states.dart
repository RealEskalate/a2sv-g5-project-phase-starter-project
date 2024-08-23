import 'package:equatable/equatable.dart';

import '../../domain/entities/product.dart';

abstract class ProductStates extends Equatable {
  final List<dynamic> data;
  const ProductStates([this.data = const <dynamic>[]]);

  @override
  List<Object?> get props => [data];
}

class InitialState extends ProductStates {}

class LoadingState extends ProductStates {}

class LoadedAllProductState extends ProductStates {
  @override
  // ignore: overridden_fields
  final List<ProductEntity> data;
  const LoadedAllProductState({required this.data}) : super(data);
}

class LoadedSingleProductState extends ProductStates {
  final ProductEntity productEntity;
  const LoadedSingleProductState({required this.productEntity});
}

class ErrorState extends ProductStates {
  final String message;
  ErrorState({required this.message}) : super([message]);
}

class SuccessfullState extends ProductStates {
  final String message;
  SuccessfullState({required this.message}) : super([message]);
}
