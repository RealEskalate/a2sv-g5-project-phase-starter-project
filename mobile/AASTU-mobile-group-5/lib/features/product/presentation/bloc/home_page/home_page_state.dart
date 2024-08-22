part of 'home_page_bloc.dart';

sealed class HomePageState extends Equatable {
  const HomePageState();
  
  @override
  List<Object> get props => [];
}

final class HomePageInitialState extends HomePageState {}

final class HomePageLoadingState extends HomePageState {}

final class HomePageLoadedState extends HomePageState {
  final List<Product> products;

  const HomePageLoadedState(this.products);

  @override
  List<Object> get props => [products];
}

final class HomePageErrorState extends HomePageState {
  final String message;

  const HomePageErrorState(this.message);

  @override
  List<Object> get props => [message];
}
