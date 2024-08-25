import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/core/constants/constants.dart';
import 'package:ecommerce_app/core/errors/failures/failure.dart';
import 'package:ecommerce_app/core/validator/validator.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  late InputDataValidator inputDataValidator;
  setUp(() {
    inputDataValidator = InputDataValidator();
  });

  group('Price validatin', () {
    test('should return appropriate data if the result is valid', () {
      final result = inputDataValidator.checkPrice('10');

      expect(result, const Right(true));
    });

    test('should return appropriate data if the result is valid', () {
      final result = inputDataValidator.checkPrice('');

      expect(result, const Right(true));
    });

    test('Should return invalid charter included when the inputs are not valid',
        () {
      final result = inputDataValidator.checkPrice('1hc');
      expect(
          result,
          Left(InvalidInputFailure(
              AppData.getMessage(AppData.invalidPriceCharacter))));
    });

    test(
        'Should return negative number cannot be price when negative numbers are inserted',
        () {
      final result = inputDataValidator.checkPrice('-10');
      expect(result,
          Left(InvalidInputFailure(AppData.getMessage(AppData.negativePrice))));
    });
  });

  group('name validator', () {
    test('Should return true when the number is valid', () {
      final result = inputDataValidator.checkNameOrCatagory('Chera Mihiretu');

      expect(result, const Right(true));
    });

    test('Should return true when the number is empty', () {
      final result = inputDataValidator.checkNameOrCatagory('Chera Mihiretu');

      expect(result, const Right(true));
    });
    test('Should return input failure when the name is not correct', () {
      final result = inputDataValidator.checkNameOrCatagory('Chera&adf');

      expect(result,
          Left(InvalidInputFailure(AppData.getMessage(AppData.invalidName))));
    });
  });

  group('checkThis test', () {
    group('email test', () {
      test('Should return true when email is fine', () {
        final result = inputDataValidator.checkThis(
            InputDataValidator.email, 'chera@gmail.com');
        expect(result, const Right(true));
      });

      test('Should return true when email is fine', () {
        final result = inputDataValidator.checkThis(
            InputDataValidator.email, 'chera@gmai');
        expect(
            result,
            Left(
                InvalidInputFailure(AppData.getMessage(AppData.invalidEmail))));
      });
    });
  });
}
