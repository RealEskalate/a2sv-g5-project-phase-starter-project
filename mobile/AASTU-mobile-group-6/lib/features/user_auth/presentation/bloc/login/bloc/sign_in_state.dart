import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_access.dart';
import 'package:equatable/equatable.dart';

import '../../../../data/models/user_model.dart';



// part of 'home_bloc.dart';

abstract class LoginState extends Equatable {
  const LoginState();
}

class LoginLoading extends LoginState {
  @override
  List<Object> get props => [];
}

class LoginLoaded extends LoginState {
    final String message;

  const LoginLoaded(this.message);

  @override
  List<Object> get props => [message];
}

class LoginFailure extends LoginState {
  final String message;

  LoginFailure(this.message);

  @override
  List<Object> get props => [message];
}

