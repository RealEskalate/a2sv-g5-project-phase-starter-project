part of 'sign_up_page_bloc.dart';

sealed class SignUpPageState extends Equatable {
  const SignUpPageState();
  
  @override
  List<Object> get props => [];
}

final class SignUpPageInitial extends SignUpPageState {}

final class SignUpPageLoading extends SignUpPageState {}

final class SignUpPageSuccess extends SignUpPageState {
  final User user;

  const SignUpPageSuccess({required this.user});

  @override
  List<Object> get props => [user];
}

final class SignUpPageFailure extends SignUpPageState {
  final String error;

  const SignUpPageFailure({required this.error});

  @override
  List<Object> get props => [error];
}