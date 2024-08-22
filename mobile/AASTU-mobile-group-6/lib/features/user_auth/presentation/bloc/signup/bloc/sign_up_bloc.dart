import 'package:bloc/bloc.dart';

import 'package:ecommerce_app_ca_tdd/core/network/network_info.dart';

import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/usecases/login_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/usecases/signup_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_state.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_state.dart';

import 'package:get/get.dart';
import 'package:http/http.dart' as http;

import 'package:internet_connection_checker/internet_connection_checker.dart';

class SignUpBloc extends Bloc<SignUpEvent, SignUpState>  {
  final client = http.Client;
  RegisterUser registerUser;

  SignUpBloc(this.registerUser) : super(SignUpLoading()) {
    on<RegisterUserEvent>((event, emit) async{
      emit(SignUpLoading());  
      var result = await registerUser(event.user);
      result.fold((l) => emit(SignUpFailure(l.message)), (r) => emit(SignUpLoaded(r)));

    });
    
  }
}