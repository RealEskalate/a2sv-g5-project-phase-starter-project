import 'package:bloc/bloc.dart';

import 'package:ecommerce_app_ca_tdd/core/network/network_info.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/get_user_info.dart';

import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/usecases/login_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/usecases/signup_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_state.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_state.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_state.dart';

import 'package:get/get.dart';
import 'package:http/http.dart' as http;

import 'package:internet_connection_checker/internet_connection_checker.dart';

class GetUserBloc extends Bloc<GetUserEvent, GetUserState>  {
  
  GetUserInfo getUserInfo;

  GetUserBloc(this.getUserInfo) : super(GetUserLoading()) {
    on<GetUserInfoEvent>((event, emit) async{
      emit(GetUserLoading());  
      var result = await getUserInfo(NoParams());
      result.fold((l) => emit(GetUserFailure(l.message)), (r) => emit(GetUserLoaded(r)));

    });
    
  }
}