import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:product_6/core/util/input_converter.dart';

void main() {
  late InputConverter inputConverter;
  setUp(
    () {
      inputConverter = InputConverter();
    },
  );

  group(
    'stringToUnsignedDouble',
    () {
      test(
          'should return unsigned double when string represents an unsigned double',
          () {
        // arrange
        const str = '123.1';

        // act
        final result = inputConverter.stringToUnsignedDouble(str);

        // result

        expect(result, const Right(123.1));
      });

      test(
          'should return Invalid Input Execption when non double string value passed',
          () {
        // arrange
        const str = 'abs';

        // act
        final result = inputConverter.stringToUnsignedDouble(str);

        // result

        expect(result, const Left(InvalidInputFailure()));
      });

      test('should return Falure  when the string is negative value', () {
        // arrange
        const str = '-123.1';

        // act
        final result = inputConverter.stringToUnsignedDouble(str);

        // result

        expect(result, const Left(InvalidInputFailure()));
      });
    },
  );
}
