

import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';

import 'package:equatable/equatable.dart';

sealed class SearchState extends Equatable {
  const SearchState();
  
  @override
  List<Object> get props => [];
}

final class SearchInitial extends SearchState {}



final class LoadingState extends SearchState {

}

final class LoadedState extends SearchState {
  final List<ProductModel> data;
  const LoadedState(this.data);

  @override
  List<Object> get props => [data];

}


final class FailedState extends SearchState {
  final String message;
  const FailedState(this.message);
}