part of 'search_page_bloc.dart';

abstract class SearchPageEvent extends Equatable {
  const SearchPageEvent();

  @override
  List<Object> get props => [];
}

class FetchAllProductsSearchEvent extends SearchPageEvent {}

class SearchProductsEvent extends SearchPageEvent {
  final String query;

  const SearchProductsEvent({required this.query});

  @override
  List<Object> get props => [query];
}
