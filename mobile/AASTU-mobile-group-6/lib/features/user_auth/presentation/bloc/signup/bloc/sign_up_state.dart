import 'package:equatable/equatable.dart';

import '../../../../data/models/user_model.dart';

abstract class SignUpState extends Equatable {
  const SignUpState();
}



class SignUpLoading extends SignUpState {
  @override
  List<Object> get props => [];
}

class SignUpLoaded extends SignUpState {
    final String message;

  const SignUpLoaded(this.message);

  @override
  List<Object> get props => [message];
}

class SignUpFailure extends SignUpState {
  final String error;

  SignUpFailure(this.error);

  @override
  List<Object> get props => [error];
}