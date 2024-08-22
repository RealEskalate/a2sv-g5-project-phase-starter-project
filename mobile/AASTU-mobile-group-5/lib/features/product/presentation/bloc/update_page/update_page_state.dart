part of 'update_page_bloc.dart';

abstract class UpdatePageState extends Equatable {
  const UpdatePageState();

  @override
  List<Object> get props => [];
}

class UpdatePageInitialState extends UpdatePageState {}

class UpdatePageSubmittingState extends UpdatePageState {}

class UpdatePageSubmittedState extends UpdatePageState {
  final ProductModel product;

  const UpdatePageSubmittedState(this.product);

  @override
  List<Object> get props => [product];
}

class UpdatePageDeletedState extends UpdatePageState {}

class UpdatePageErrorState extends UpdatePageState {
  final String message;

  const UpdatePageErrorState(this.message);

  @override
  List<Object> get props => [message];
}
