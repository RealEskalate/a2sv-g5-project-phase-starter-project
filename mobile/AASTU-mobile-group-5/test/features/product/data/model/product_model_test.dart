import 'package:flutter_test/flutter_test.dart';
import 'package:task_9/features/product/data/models/product_model.dart';
import 'package:task_9/features/product/domain/entities/product.dart';

void main() {
  const tProductModel = ProductModel(
    id: '1',
    name: 'Test Product',
    description: 'This is a test product',
    price: 200.0,
    imageUrl: 'assets/images/boots.jpg',
  );

  final tJson = {
    'id': '1',
    'name': 'Test Product',
    'description': 'This is a test product',
    'price': 200.0,
    'imageUrl': 'assets/images/boots.jpg',
  };

  test(
    'should be a subclass of Product entity',
    () async {
      // assert
      expect(tProductModel, isA<Product>());
    },
  );

  test(
    'should correctly convert from JSON to ProductModel',
    () async {
      final Map<String, dynamic> jsonMap = tJson;
      final result = ProductModel.fromJson(jsonMap);
      expect(result, tProductModel);
    },
  );

  test(
    'should correctly convert to JSON from ProductModel',
    () async {
      final result = tProductModel.toJson();
      expect(result, tJson);
    },
  );
}
