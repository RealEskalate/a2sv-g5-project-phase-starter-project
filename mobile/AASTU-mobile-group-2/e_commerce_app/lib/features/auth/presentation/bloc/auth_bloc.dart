import 'dart:async';
import 'package:e_commerce_app/features/auth/domain/usecase/get_user.dart';
import 'package:e_commerce_app/features/auth/domain/usecase/login.dart';
import 'package:e_commerce_app/features/auth/presentation/bloc/auth_event.dart';
import 'package:e_commerce_app/features/auth/presentation/bloc/auth_state.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/usecase/signup.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  Login login;
  SignUp signUp;
  GetUser getUser;

  AuthBloc({required this.login, required this.signUp,required this.getUser}) : super(AuthInitial()) {
    on<LoginEvent>(_onLogin);
    on<SignUpEvent>(_onSignUp);
    on<GetUserEvent>(_onGetUser);
  }

  FutureOr<void> _onLogin(LoginEvent event, Emitter<AuthState> emit) async {
    emit(LoginLoading());
    final result =
        await login.execute(email: event.email, password: event.password);

    print(result);
    result.fold((failure) => emit(AuthFailure()),
        (token) => emit(LoginSuccess(token: token)));
  }

  FutureOr<void> _onSignUp(SignUpEvent event, Emitter<AuthState> emit) async {
    emit(SignUpLoading());
    final result = await signUp.execute(
        name: event.name, email: event.email, password: event.password);
    print(result);

    result.fold((failure) => emit(AuthFailure()),
        (user) => emit(SignUpSuccess(user: user)));
  }

  FutureOr<void> _onGetUser(GetUserEvent event, Emitter<AuthState> emit) {
    emit(GetUserLoading());
    getUser.execute().then((value) {
      value.fold((failure) => emit(AuthFailure()),
          (name) => emit(GetUserSuccess(name: name)));
    });
  }
}
