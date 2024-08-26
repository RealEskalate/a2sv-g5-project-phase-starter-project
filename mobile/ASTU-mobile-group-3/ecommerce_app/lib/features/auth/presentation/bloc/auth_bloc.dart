// ignore: depend_on_referenced_packages
import 'dart:developer';

import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:http/http.dart';

import '../../domain/entities/user_entity.dart';
import '../../domain/usecases/get_me_usecase.dart';
import '../../domain/usecases/log_in_usecase.dart';
import '../../domain/usecases/log_out_usecase.dart';
import '../../domain/usecases/sign_up_usecase.dart';

part 'auth_event.dart';
part 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final SignUpUsecase signUpUsecase;
  final LogInUsecase logInUsecase;
  final LogOutUsecase logOutUsecase;
  final GetMeUsecase getMeUsecase;
  AuthBloc(
      {required this.signUpUsecase,
      required this.logInUsecase,
      required this.logOutUsecase,
      required this.getMeUsecase})
      : super(AuthInitial()) {
    on<LogInEvent>((event, emit) async {
      emit(AuthLoadingstate());
      UserEntity myEnity = UserEntity(
        name: '',
        email: event.email,
        password: event.password,
        id: '',
        v: 0,
      );
      final result = await logInUsecase.execute(myEnity);
      result.fold((failure) {
        emit(LoginErrorState(message: failure.message));
      }, (data) {
        emit(LogInSuccessState());
      });
    });
    on<SignUpEvent>((event, emit) async {
      emit(AuthLoadingstate());
      UserEntity myEnity = UserEntity(
        name: event.name,
        email: event.email,
        password: event.password,
        id: '',
        v: 0,
      );
      final result = await signUpUsecase.execute(myEnity);

      result.fold((failure) {
        emit(SignupErrorState(message: failure.message));
      }, (data) {
        emit(RegisterSuccessState());
      });
    });

    on<LogOutEvent>((event, emit) async {
      final result = await logOutUsecase.execute();
      result.fold((failure) {
        emit(LogoutFailedState(message: failure.message));
      }, (data) {
        emit(LogoutSuccess());
      });
    });

    on<GetMeEvent>((event, emit) async {
      emit(AuthLoadingstate());

      final result = await getMeUsecase.execute();
      result.fold((failure) {
        emit(LogoutFailedState(message: failure.message));
      }, (data) {
        emit(GetMeSuccessState(user: data));
      });
    });
  }
}
