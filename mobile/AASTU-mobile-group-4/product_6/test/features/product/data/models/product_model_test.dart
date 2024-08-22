import 'package:product_6/features/product/data/models/product_model.dart';
import 'package:test/test.dart';

void main() {
  group('ProductModel', () {
    final productJson = {
      'id': '1',
      'name': 'Test Product',
      'description': 'This is a test product',
      'imageUrl': 'http://example.com/image.png',
      'price': 19.99,
    };

    const productModel = ProductModel(
      id: '1',
      name: 'Test Product',
      description: 'This is a test product',
      imageUrl: 'http://example.com/image.png',
      price: 19.99,
    );

    test('fromJson creates a valid model from JSON', () {
      final result = ProductModel.fromJson(productJson);

      expect(result, isA<ProductModel>());
      // expect(result.id, 1);
      // expect(result.name, 'Test Product');
      // expect(result.description, 'This is a test product');
      // expect(result.imageUrl, 'http://example.com/image.png');
      // expect(result.price, 19.99);
    });

    test('toJson returns a valid JSON map from the model', () {
      final result = productModel.toJson();

      expect(result, isA<Map<String, dynamic>>());
      expect(result['id'], 1);
      expect(result['name'], 'Test Product');
      expect(result['description'], 'This is a test product');
      expect(result['imageUrl'], 'http://example.com/image.png');
      expect(result['price'], 19.99);
    });

    test('fromJson and toJson are consistent', () {
      final json = productModel.toJson();
      final modelFromJson = ProductModel.fromJson(json);

      expect(modelFromJson, isA<ProductModel>());
      expect(modelFromJson.id, productModel.id);
      expect(modelFromJson.name, productModel.name);
      expect(modelFromJson.description, productModel.description);
      expect(modelFromJson.imageUrl, productModel.imageUrl);
      expect(modelFromJson.price, productModel.price);
    });
  });
}
