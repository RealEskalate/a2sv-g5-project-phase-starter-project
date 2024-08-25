import 'package:bloc_test/bloc_test.dart';
import 'package:ecommerce_app/core/validator/validator.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/cubit/input_validation_cubit.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  late InputValidationCubit inputValidationCubit;
  late InputDataValidator inputDataValidator;
  setUp(() {
    inputDataValidator = InputDataValidator();
    inputValidationCubit = InputValidationCubit(inputDataValidator);
  });

  group('Cubit test on input', () {
    test('Initial state should be set', () {
      expect(inputValidationCubit.state, isA<InputValidationInitial>());
    });

    blocTest(
      'Should return correspond value as true',
      build: () {
        return inputValidationCubit;
      },
      act: (bloc) => bloc.checkChanges(['Name', 'Ch']),
      expect: () => [
        InputValidatedState(name: true, catagory: true, price: true),
      ],
    );

    blocTest(
      'Should return correspond value as when name is valid',
      build: () {
        return inputValidationCubit;
      },
      act: (bloc) => bloc.checkChanges(['Name', 'Chera']),
      expect: () => [
        InputValidatedState(name: true, catagory: true, price: true),
      ],
    );
    blocTest(
      'Should return correspond value as when name is not valid',
      build: () {
        return inputValidationCubit;
      },
      act: (bloc) => bloc.checkChanges(['Name', 'Chera 87']),
      expect: () => [
        InputValidatedState(name: false, catagory: true, price: true),
      ],
    );

    blocTest(
      'Should return correspond value as when price is not valid',
      build: () {
        return inputValidationCubit;
      },
      act: (bloc) => bloc.checkChanges(['Price', 'Ch']),
      expect: () => [
        InputValidatedState(name: true, catagory: true, price: false),
      ],
    );

    blocTest(
      'Should return correspond value as when price is  valid',
      build: () {
        return inputValidationCubit;
      },
      act: (bloc) => bloc.checkChanges(['Price', '10']),
      expect: () => [
        InputValidatedState(name: true, catagory: true, price: true),
      ],
    );

    blocTest(
      'Should nothing when unkown input field is called',
      build: () {
        return inputValidationCubit;
      },
      act: (bloc) => bloc.checkChanges(['Tola', '10']),
      expect: () => [],
    );
  });
}
