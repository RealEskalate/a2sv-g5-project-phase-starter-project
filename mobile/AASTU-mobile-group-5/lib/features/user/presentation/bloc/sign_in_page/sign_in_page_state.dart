part of 'sign_in_page_bloc.dart';

sealed class SignInPageState extends Equatable {
  const SignInPageState();
  
  @override
  List<Object> get props => [];
}

final class SignInPageInitial extends SignInPageState {}

final class SignInPageLoading extends SignInPageState {}

final class SignInPageSuccess extends SignInPageState {
  final String user;

  const SignInPageSuccess({required this.user});

  @override
  List<Object> get props => [user];
}

final class SignInPageFailure extends SignInPageState {
  final String error;

  const SignInPageFailure({required this.error});

  @override
  List<Object> get props => [error];
}
