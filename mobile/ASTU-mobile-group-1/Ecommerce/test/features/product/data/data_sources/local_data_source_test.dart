// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
// import 'package:product_6/core/error/exception.dart';
// import 'package:product_6/features/product/data/data_sources/local_data_source.dart';
// import 'package:product_6/features/product/data/models/product_model.dart';

// import '../../../../helpers/json_reader.dart';
// import '../../../../helpers/test_helper.mocks.dart';

// void main() {
//   late MockSharedPreferences mockSharedPreferences;
//   late ProductLocalDataSourceImpl productLocalDataSourceImpl;

//   setUp(() {
//     mockSharedPreferences = MockSharedPreferences();
//     productLocalDataSourceImpl =
//         ProductLocalDataSourceImpl(pref: mockSharedPreferences);
//   });

//   group('getCachedProducts', () {
//     test('should return a list of product model', () async {
//       // arrange
//       when(mockSharedPreferences.getStringList('cachedProduct'))
//           .thenReturn([readJson('helpers/dummy_data/dummy_product_data.json')]);

//       // act

//       final result = await productLocalDataSourceImpl.getCachedProducts();

//       // assert
//       expect(result, isA<List<ProductModel>>());
//     });

//     test('should throw a CacheException when there is no cached value',
//         () async {
//       // arrange
//       when(mockSharedPreferences.getStringList('cachedProduct'))
//           .thenReturn(null);

//       // act
//       final call = productLocalDataSourceImpl.getCachedProducts;

//       // assert
//       expect(() => call(), throwsA(isA<CacheException>()));
//     });
//   });
// }
