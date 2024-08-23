


import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../Domain/usecase/ecommerce_usecase.dart';
import 'login_user_states_event.dart';
import 'login_user_states_state.dart';

class LoginUserStatesBloc  extends Bloc<LoginUserStatesEvent,LoginUserStates>{
  final EcommerceUsecase ecommerceUsecase;
  LoginUserStatesBloc({required this.ecommerceUsecase}) : super(LeftUserStates()){
    

    on<LogedOutUserStatesEvent>((event, emit) async{
      final name = await ecommerceUsecase.deleteToken('name');
      final email = await ecommerceUsecase.deleteToken('email');
      final token = await ecommerceUsecase.deleteToken('key');
   
      if (name && email && token) {
        emit(LogedOutUserStates());
      }
      emit(LeftUserStates());
    });

    on<ProfileDetail>((event, emit) async{
      final result = await ecommerceUsecase.getName('name');
      final emial = await ecommerceUsecase.getName('email');
      emit(ProfileDetailState(name: result,email: emial));
    });
  }

}