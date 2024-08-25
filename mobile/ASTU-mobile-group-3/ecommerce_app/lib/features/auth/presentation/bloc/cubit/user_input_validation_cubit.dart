// ignore: depend_on_referenced_packages
import 'package:bloc/bloc.dart';
import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../../core/constants/constants.dart';
import '../../../../../core/errors/failures/failure.dart';
import '../../../../../core/validator/validator.dart';

part 'user_input_validation_state.dart';

class UserInputValidationCubit extends Cubit<UserInputValidationState> {
  final InputDataValidator inputDataValidator;
  UserInputValidationCubit({required this.inputDataValidator})
      : super(UserInputValidationInitial());

  void checkWith(String from, String type, String val, [String? conf]) {
    if (type == InputDataValidator.confirmPass && conf == null) {
      checkWith(from, type, val, state.passwordContent);
      return;
    }

    Map<String, String> correspond = {
      InputDataValidator.name: state.name,
      InputDataValidator.email: state.email,
      InputDataValidator.password: state.password,
      InputDataValidator.confirmPass: state.confirmPassword,
      InputDataValidator.checkBox: state.checkbox,
      InputDataValidator.pc: state.passwordContent
    };

    Either<Failure, bool> result;
    if (from == AppData.login) {
      result = inputDataValidator.checkThis(type, val);
      result.fold((failure) {
        correspond[type] = AppData.strNotValidated;
      }, (data) {
        correspond[type] = AppData.strValidated;
      });

      emit(
        LoginUserInputValidated(
          email: correspond[InputDataValidator.email]!,
          password: correspond[InputDataValidator.password]!,
          checkbox: correspond[InputDataValidator.checkBox]!,
        ),
      );
    } else {
      if (type != InputDataValidator.confirmPass) {
        result = inputDataValidator.checkThis(type, val);
      } else {
        result = inputDataValidator.checkThis(type, val, conf);
      }

      result = inputDataValidator.checkThis(type, val);

      result.fold((failure) {
        correspond[type] = AppData.strNotValidated;
      }, (data) {
        correspond[type] = AppData.strValidated;
      });

      emit(SignupUserInputValidated(
          name: correspond[InputDataValidator.name]!,
          email: correspond[InputDataValidator.email]!,
          password: correspond[InputDataValidator.password]!,
          checkbox: correspond[InputDataValidator.checkBox]!,
          confirmPassword: correspond[InputDataValidator.confirmPass]!,
          passwordContent: correspond[InputDataValidator.pc]!));
    }
  }

  void reset() {
    emit(UserInputValidationInitial());
  }

  void changeCheckbox(String from, bool val) {
    Map<String, String> correspond = {
      InputDataValidator.name: state.name,
      InputDataValidator.email: state.email,
      InputDataValidator.password: state.password,
      InputDataValidator.confirmPass: state.confirmPassword,
      InputDataValidator.checkBox: state.checkbox,
      InputDataValidator.pc: state.password
    };

    if (val == true) {
      correspond[InputDataValidator.checkBox] = AppData.strValidated;
    } else {
      correspond[InputDataValidator.checkBox] = AppData.strNotValidated;
    }

    if (from == AppData.signup) {
      emit(
        SignupUserInputValidated(
            name: correspond[InputDataValidator.name]!,
            email: correspond[InputDataValidator.email]!,
            password: correspond[InputDataValidator.password]!,
            checkbox: correspond[InputDataValidator.checkBox]!,
            confirmPassword: correspond[InputDataValidator.confirmPass]!,
            passwordContent: correspond[InputDataValidator.pc]!),
      );
    } else {
      emit(LoginUserInputValidated(
        email: correspond[InputDataValidator.email]!,
        password: correspond[InputDataValidator.password]!,
        checkbox: correspond[InputDataValidator.checkBox]!,
      ));
    }
  }
}
