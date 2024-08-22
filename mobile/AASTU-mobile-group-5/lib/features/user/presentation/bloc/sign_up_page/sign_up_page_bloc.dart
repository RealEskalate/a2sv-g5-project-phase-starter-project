import 'dart:async';

import 'package:bloc/bloc.dart';
import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../../core/failure/failure.dart';
import '../../../domain/entities/user.dart';
import '../../../domain/repositories/user_repository.dart';

part 'sign_up_page_event.dart';
part 'sign_up_page_state.dart';

class SignUpPageBloc extends Bloc<SignUpPageEvent, SignUpPageState> {
  final UserRepository userRepository;

  SignUpPageBloc({required this.userRepository}) : super(SignUpPageInitial()){
    on<SignUpPageButtonPressed>(_onsignupbuttonpressed);
    
  }


  FutureOr<void> _onsignupbuttonpressed(SignUpPageButtonPressed event, Emitter<SignUpPageState> emit)async {
    if (event.password != event.confirmPassword) {
        emit (const SignUpPageFailure(error: 'Passwords do not match'));
      }

      emit (SignUpPageLoading());

      final Either<Failure, User> failureOrUser = await userRepository.registerUser(
        event.email,
        event.password,
        event.name,
      );

      emit (failureOrUser.fold(
        (failure) => SignUpPageFailure(error: failure.message),
        (user) => SignUpPageSuccess(user: user),
      ));
  }
}


