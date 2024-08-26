import 'package:equatable/equatable.dart';

sealed class SearchEvent extends Equatable {
  const SearchEvent();

  @override
  List<Object> get props => [];
}

class LoadAllProductEvent extends SearchEvent {}

class SearchProductEvent extends SearchEvent {
  final String query;

  const SearchProductEvent(this.query);

  @override
  List<Object> get props => [query];
}