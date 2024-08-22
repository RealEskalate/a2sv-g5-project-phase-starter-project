

import 'package:ecommerce/features/product/data/data_sources/local_data_source.dart';
import 'package:ecommerce/features/product/data/model/product_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:shared_preferences/shared_preferences.dart';

void main() async {
  SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
  var localDataSource = localDataSourceImpl(sharedPreference: sharedPreferences);
  setUp((){});
  test("",() async {
  
  const tProductModel =  ProductModel(
      id: '123',
      name: 'Nike Shoes',
      description: 'High-quality sports shoes',
      price: 99,
      image: 'path/to/image.jpg'
    );
    var listProduct = [tProductModel];
    var all = await localDataSource.cachedproduct(listProduct);
    var data = localDataSource.getallproduct();
  });
}