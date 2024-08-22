import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/usecase/usecase.dart';
import '../../domain/entities/log_in.dart';
import '../../domain/entities/sign_up.dart';
import '../../domain/entities/user_data.dart';
import '../../domain/usecases/get_current_user_usecase.dart';
import '../../domain/usecases/log_in_usecase.dart';
import '../../domain/usecases/log_out_usecase.dart';
import '../../domain/usecases/sign_up_usecase.dart';

part 'auth_event.dart';
part 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final GetCurrentUserUsecase getCurrentUserUsecase;
  final LogInUsecase logInUsecase;
  final LogOutUsecase logOutUsecase;
  final SignUpUsecase signUpUsecase;
  AuthBloc({
    required this.getCurrentUserUsecase,
    required this.logInUsecase,
    required this.logOutUsecase,
    required this.signUpUsecase,
  }) : super(AuthInitial()) {
    on<GetCurrentUserEvent>((event, emit) async {
      emit(AuthLoadingState());

      final result = await getCurrentUserUsecase(NoParams());
      result.fold(
        (failure) => emit(AuthErrorState(message: failure.message)),
        (success) => emit(AuthUserLoaded(userEntity: success)),
      );
    });
    on<LogInEvent>((event, emit) async {
      emit(AuthLoadingState());

      final result =
          await logInUsecase(LogInParams(logInEntity: event.logInEntity));
      result.fold(
        (failure) => emit(AuthErrorState(message: failure.message)),
        (success) => emit(AuthSignedInState()),
      );
    });
    on<SignUpEvent>((event, emit) async {
      emit(AuthLoadingState());

      final result =
          await signUpUsecase(GetParams(signUpEntity: event.signUpEntity));
      result.fold(
        (failure) => emit(AuthErrorState(message: failure.message)),
        (success) => emit(AuthSignedUpState()),
      );
    });
    on<LogOutEvent>((event, emit) async {
      emit(AuthLoadingState());

      final result = await logOutUsecase(NoParams());
      result.fold(
        (failure) => emit(AuthErrorState(message: failure.message)),
        (success) => emit(AuthLogOutState()),
      );
    });
  }
}
