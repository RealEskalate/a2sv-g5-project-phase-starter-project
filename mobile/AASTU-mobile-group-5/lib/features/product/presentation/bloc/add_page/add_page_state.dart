// add_page_state.dart
part of 'add_page_bloc.dart';

sealed class AddPageState extends Equatable {
  const AddPageState();
  
  @override
  List<Object> get props => [];
}

final class AddPageInitialState extends AddPageState {}

final class AddPageSubmittingState extends AddPageState {}

final class AddPageSubmittedState extends AddPageState {
  final ProductModel product;

  const AddPageSubmittedState(this.product);

  @override
  List<Object> get props => [product];
}

final class UpdatePageSubmittedState extends AddPageState { // New state for updates
  final ProductModel product;

  const UpdatePageSubmittedState(this.product);

  @override
  List<Object> get props => [product];
}


final class AddPageErrorState extends AddPageState {
  final String message;

  const AddPageErrorState(this.message);

  @override
  List<Object> get props => [message];
}
