


import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../domain/usecase/login_usecase.dart';
import 'login_event.dart';
import 'login_state.dart';


class LoginBloc extends Bloc<LoginEvent,LoginState>{
      final LoginUseCase loginUseCase;
      LoginBloc({
        required this.loginUseCase
      }):super(LoginIntial()){

      on<LoginRequest> (
        (event,emit) async {
          final result = await loginUseCase.loginUser(event.email, event.password);
          result.fold(
            (failure) {
              emit( LoginSFuiled(message: 'try agine pless'));
            },
            (data) {
              emit(LoginSuccess(message: 'login Success'));
            }

          );
        }


      );
      }


}