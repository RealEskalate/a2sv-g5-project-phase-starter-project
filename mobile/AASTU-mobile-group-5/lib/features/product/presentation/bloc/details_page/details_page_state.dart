// details_page_state.dart
part of 'details_page_bloc.dart';

abstract class DetailsPageState extends Equatable {
  const DetailsPageState();

  @override
  List<Object> get props => [];
}

class DetailsPageInitialState extends DetailsPageState {}

class DetailsPageLoadingState extends DetailsPageState {}

class DetailsPageLoadedState extends DetailsPageState {
  final Product product;

  const DetailsPageLoadedState(this.product);

  @override
  List<Object> get props => [product];
}

class DetailsPageDeletedState extends DetailsPageState {}

class DetailsPageErrorState extends DetailsPageState {
  final String message;

  const DetailsPageErrorState(this.message);

  @override
  List<Object> get props => [message];
}
