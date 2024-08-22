part of 'home_page_bloc.dart';

sealed class HomePageEvent extends Equatable {
  const HomePageEvent();

  @override
  List<Object> get props => [];
}

class FetchAllProductsEvent extends HomePageEvent {}

class AddProductToHomePageEvent extends HomePageEvent {
  final Product product;

  const AddProductToHomePageEvent(this.product);

  @override
  List<Object> get props => [product];
}