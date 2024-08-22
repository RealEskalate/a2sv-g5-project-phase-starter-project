import 'package:bloc/bloc.dart';

import 'package:ecommerce_app_ca_tdd/core/network/network_info.dart';

import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/usecases/login_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/usecases/signup_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_state.dart';

import 'package:get/get.dart';
import 'package:http/http.dart' as http;

import 'package:internet_connection_checker/internet_connection_checker.dart';

class LoginBloc extends Bloc<LoginEvent, LoginState>  {
  final client = http.Client;
  LoginUsecase loginUsecase;

  LoginBloc(this.loginUsecase) : super(LoginLoading()) {
    on<LogUserIn>((event, emit) async{
      emit(LoginLoading());  
      var result = await loginUsecase(event.user);
      result.fold((l) => emit(LoginFailure(l.message)), (r) => emit(LoginLoaded(r)));

    });
    
  }
}
