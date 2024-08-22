part of 'search_page_bloc.dart';

sealed class SearchPageState extends Equatable {
  const SearchPageState();
  
  @override
  List<Object> get props => [];
}

final class SearchPageInitialState extends SearchPageState {}

final class SearchPageLoadingState extends SearchPageState{}

final class SearchPageLoadedState extends SearchPageState {
  final List<Product> products;

  const SearchPageLoadedState(this.products);

  @override
  List<Object> get props => [products];
}

final class SearchPageErrorState extends SearchPageState {
  final String message;

  const SearchPageErrorState(this.message);

  @override
  List<Object> get props => [message];
}

