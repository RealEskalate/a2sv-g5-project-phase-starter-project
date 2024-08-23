import 'package:dartz/dartz.dart';

import '../constants/constants.dart';
import '../errors/failures/failure.dart';

class InputDataValidator {
  static const String name = 'Name';
  static const String email = 'Email';
  static const String password = 'Password';
  static const String price = 'Price';
  static const String catagory = 'Catagory';
  static const String confirmPass = 'Confirm Password';
  static const String checkBox = 'CheckBox';
  static const String pc = 'PC';

  Either<Failure, bool> checkPrice(String input) {
    if (input.isEmpty) return const Right(true);
    int value = 0;
    try {
      value = int.parse(input);
      if (value < 0) {
        return Left(
            InvalidInputFailure(AppData.getMessage(AppData.negativePrice)));
      }
      return const Right(true);
    } on FormatException {
      return Left(InvalidInputFailure(
          AppData.getMessage(AppData.invalidPriceCharacter)));
    }
  }

  Either<Failure, bool> checkNameOrCatagory(String input) {
    if (input.isEmpty) {
      return Left(InvalidInputFailure(AppData.getMessage(AppData.invalidName)));
    }

    /// The name validator will goo here
    ///
    final RegExp nameRegExp = RegExp(r'^[a-zA-Z\s]+$');
    if (nameRegExp.hasMatch(input)) {
      return const Right(true);
    } else {
      return Left(InvalidInputFailure(AppData.getMessage(AppData.invalidName)));
    }
  }

  Either<Failure, bool> checkPassword(String input) {
    if (input.length >= 6) {
      return const Right(true);
    } else {
      return Left(
          InvalidInputFailure(AppData.getMessage(AppData.invalidPassword)));
    }
  }

  Either<Failure, bool> checkConfirmPassword(
      String password, String confirmpass) {
    if (password == confirmpass) {
      return const Right(true);
    } else {
      return Left(
          InvalidInputFailure(AppData.getMessage(AppData.confirmPassword)));
    }
  }

  Either<Failure, bool> checkEmail(String email) {
    final RegExp emailRegExp = RegExp(
      r'^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$',
      caseSensitive: false,
      multiLine: false,
    );

    if (emailRegExp.hasMatch(email)) {
      return const Right(true);
    } else {
      return Left(
          InvalidInputFailure(AppData.getMessage(AppData.invalidEmail)));
    }
  }

  Either<Failure, bool> checkThis(String type, String val, [String? moreVal]) {
    Map<String, dynamic> map = {
      name: checkNameOrCatagory,
      catagory: checkNameOrCatagory,
      password: checkPassword,
      email: checkEmail,
      confirmPass: checkConfirmPassword,
      price: checkPrice,
    };
    if (type == confirmPass && moreVal == null) {
      return const Left(InvalidInputFailure(''));
    }
    if (moreVal != null) {
      return map[type](val, moreVal);
    } else {
      return map[type](val);
    }
  }
}
