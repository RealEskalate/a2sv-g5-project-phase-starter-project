

import 'package:email_validator/email_validator.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../domain/usecase/login_usecase.dart';
import 'login_registration_event.dart';
import 'login_registration_state.dart';


class LoginRegistrationBloc extends Bloc<LoginRegistrationEvent,LoginRegistrationState> {
  final LoginUseCase loginUseCase;
  LoginRegistrationBloc({required this.loginUseCase}):super(const InisialState()){
    String password = '';
    String confirmPassword = '';
    String email = '';
    String fullName = '';
    String newPassword = '';
    String newEmail = '';
    bool terms = false;

    on<OnInputChangeEvent>(
      (event,emit){
        if (event.type == 'email') {
        
          email = event.email;

        } else if (event.type == 'password') {
          password = event.password;
        } else if (event.type == 'confirmPassword') {
          confirmPassword = event.confirmPassword;
        } else if (event.type == 'fullName') {
       
          fullName = event.fullName;
        } else if (event.type == 'newEmail') {
  
          newEmail = event.newEmail;
        } else if (event.type == 'newPassword') {
          newPassword = event.newPassword;
        } else if(event.type == 'terms'){
         
          terms = event.terms;
          
        }
        emit(const InisialState());
        emit(OnInputChange(email: email, password: password, confirmPassword: confirmPassword, fullName: fullName, newEmail: newEmail, newPassword: newPassword, terms: terms));  
      }
    );


    on<LoginButtonPressed> (
      (event,emit)async {
        final bool emailChecker = EmailValidator.validate(email);
        
        if (email == '') {
          emit(const OnErrorState(error: 'Email is required',email: true));
        } else if (!emailChecker) {
            emit(const OnErrorState(error: 'Invalid Email',email: true));
          }
        else if (password == '') {
          emit(const OnErrorState(error: 'Password is required',password: true));
        } 
        else {
          emit(const OnLoading());
          final result = await loginUseCase.loginUser(email, password);
          
          result.fold(
            (failure) {
         
              emit( OnErrorState(error: failure.message ));
            },
            (data) {
              email = '';
              password = '';
      
              emit(LoginSuccess(email: data.email, name: data.name,));
            }
          );
          
        }
      }
    );


    on<RegisterButtonPressed>(
      (event,emit) async{
        
        final bool emailChecker = EmailValidator.validate(newEmail);
        if (fullName == '') {
          emit(const OnErrorState(error: 'Full name is required',fullName: true));
        }else if (newEmail == '') {         
          emit(const OnErrorState(error: 'email is required',newEmail: true));
        } else if (!emailChecker) {
            emit(const OnErrorState(error: 'Invalid Email',newEmail: true));
          }
        else if (newPassword == '') {
          emit(const OnErrorState(error: 'Password is required',newPassword: true));
        } else if (confirmPassword == '') {
          emit(const OnErrorState(error: 'Confirm Password is required',confirmPassword: true));
        } else if (newPassword != confirmPassword) {
          emit(const OnErrorState(error: 'Password does not match',newPassword: true,confirmPassword: true));
        } else if(terms == false){
          emit(const OnErrorState(error: 'Please accept terms and conditions',terms: true));
        } 
        else {
         
          emit(const OnLoading());
          final result = await loginUseCase.registerUser(newEmail, newPassword,fullName);
          result.fold(
            (failure) {
              
              emit( OnErrorState(error: failure.message ));
            },
            (data) {
            
              emit(RegistrationSuccess(success: data));
            }
          );
          
        }
      }
    );
  }

  
  
}