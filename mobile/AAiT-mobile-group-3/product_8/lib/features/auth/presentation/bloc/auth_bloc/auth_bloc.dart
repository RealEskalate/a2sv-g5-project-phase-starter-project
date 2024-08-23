import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../../../core/usecase/usecase.dart';
import '../../../domain/entities/sign_in_user_entitiy.dart';
import '../../../domain/entities/sign_up_user_entitiy.dart';
import '../../../domain/entities/user_data_entity.dart';
import '../../../domain/use_case/get_user_use_case.dart';
import '../../../domain/use_case/log_out_use_case.dart';
import '../../../domain/use_case/sign_in_use_case.dart';
import '../../../domain/use_case/sign_up_use_case.dart';

part 'auth_event.dart';
part 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final SignInUseCase _signInUseCase;
  final SignUpUseCase _signUpUseCase;
  final GetUserUseCase _getUserUsecase;
  final LogOutUseCase _logoutUsecase;

  AuthBloc(this._signInUseCase, this._signUpUseCase, this._getUserUsecase,
      this._logoutUsecase)
      : super(AuthInitial()) {
    on<SigninEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await _signInUseCase(
          SignInParams(signInUserEntitiy: event.signInUserEntitiy));
      result.fold((failure) {
        print(failure);
        emit(AuthError(message: 'Invalid input'));
      }, (userData) => emit(AuthSuccess()));
    });

    on<SignupEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await _signUpUseCase(
          SignUpParams(signUpUserEntitiy: event.signUpUserEntitiy));
      result.fold((failure) {
        print(failure);
        emit(AuthError(message: 'Invalid input'));
      }, (userData) => emit(AuthRegisterSuccess()));
    });

    on<GetUserEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await _getUserUsecase(NoParams());
      result.fold((failure) => emit(AuthError(message: 'Invalid input')),
          (userData) => emit(AuthAuthenticated(userDataEntity: userData)));
    });

    on<LogoutEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await _logoutUsecase(NoParams());
      result.fold((failure) => emit(AuthError(message: 'Invalid input')),
          (userData) => emit(AuthLoggedOut()));
    });
  }
}
