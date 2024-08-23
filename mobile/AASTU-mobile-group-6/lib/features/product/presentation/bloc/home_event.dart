import 'package:equatable/equatable.dart';


abstract class HomeEvent extends Equatable {
  const HomeEvent();
}

class GetProductsEvent extends HomeEvent {
  @override
  List<Object> get props => [];
}

class SearchProductsEvent extends HomeEvent {
  final String searchTerm;

  SearchProductsEvent(this.searchTerm);

  @override
  List<Object> get props => [searchTerm];
}
