import 'dart:async';

import 'package:bloc/bloc.dart';
import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../../core/failure/failure.dart';
import '../../../domain/repositories/user_repository.dart';

part 'sign_in_page_event.dart';
part 'sign_in_page_state.dart';



class SignInPageBloc extends Bloc<SignInPageEvent, SignInPageState> {
  final UserRepository userRepository;

  SignInPageBloc({required this.userRepository}) : super(SignInPageInitial()){
    on<SignInPageButtonPressed>(_onsignInbuttonpressed);
    
  }

  FutureOr<void> _onsignInbuttonpressed(SignInPageButtonPressed event, Emitter<SignInPageState> emit)async {
    emit(SignInPageLoading());

      final Either<Failure, String> failureOrUser = await userRepository.loginUser(
        event.email,
        event.password,
      );

      emit( failureOrUser.fold(
        (failure) => SignInPageFailure(error: failure.message),
        (user) => SignInPageSuccess(user: user),
      ));
  }
}