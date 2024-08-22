// Mocks generated by Mockito 5.4.4 from annotations
// in task_9/test/features/product/domain/use_case/delete_product_test.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i4;

import 'package:dartz/dartz.dart' as _i2;
import 'package:mockito/mockito.dart' as _i1;
import 'package:task_9/core/failure/failure.dart' as _i5;
import 'package:task_9/features/product/data/models/product_model.dart' as _i7;
import 'package:task_9/features/product/domain/entities/product.dart' as _i6;
import 'package:task_9/features/product/domain/repository/product_repository.dart'
    as _i3;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: deprecated_member_use
// ignore_for_file: deprecated_member_use_from_same_package
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

class _FakeEither_0<L, R> extends _i1.SmartFake implements _i2.Either<L, R> {
  _FakeEither_0(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

/// A class which mocks [ProductRepository].
///
/// See the documentation for Mockito's code generation for more information.
class MockProductRepository extends _i1.Mock implements _i3.ProductRepository {
  MockProductRepository() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i4.Future<_i2.Either<_i5.Failure, List<_i6.Product>>> getAllProducts() =>
      (super.noSuchMethod(
        Invocation.method(
          #getAllProducts,
          [],
        ),
        returnValue:
            _i4.Future<_i2.Either<_i5.Failure, List<_i6.Product>>>.value(
                _FakeEither_0<_i5.Failure, List<_i6.Product>>(
          this,
          Invocation.method(
            #getAllProducts,
            [],
          ),
        )),
      ) as _i4.Future<_i2.Either<_i5.Failure, List<_i6.Product>>>);

  @override
  _i4.Future<_i2.Either<_i5.Failure, _i6.Product>> getProductById(String? id) =>
      (super.noSuchMethod(
        Invocation.method(
          #getProductById,
          [id],
        ),
        returnValue: _i4.Future<_i2.Either<_i5.Failure, _i6.Product>>.value(
            _FakeEither_0<_i5.Failure, _i6.Product>(
          this,
          Invocation.method(
            #getProductById,
            [id],
          ),
        )),
      ) as _i4.Future<_i2.Either<_i5.Failure, _i6.Product>>);

  @override
  _i4.Future<_i2.Either<_i5.Failure, _i7.ProductModel>> addProduct(
    _i7.ProductModel? product,
    String? imagePath,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #addProduct,
          [
            product,
            imagePath,
          ],
        ),
        returnValue:
            _i4.Future<_i2.Either<_i5.Failure, _i7.ProductModel>>.value(
                _FakeEither_0<_i5.Failure, _i7.ProductModel>(
          this,
          Invocation.method(
            #addProduct,
            [
              product,
              imagePath,
            ],
          ),
        )),
      ) as _i4.Future<_i2.Either<_i5.Failure, _i7.ProductModel>>);

  @override
  _i4.Future<_i2.Either<_i5.Failure, _i7.ProductModel>> updateProduct(
          _i7.ProductModel? product) =>
      (super.noSuchMethod(
        Invocation.method(
          #updateProduct,
          [product],
        ),
        returnValue:
            _i4.Future<_i2.Either<_i5.Failure, _i7.ProductModel>>.value(
                _FakeEither_0<_i5.Failure, _i7.ProductModel>(
          this,
          Invocation.method(
            #updateProduct,
            [product],
          ),
        )),
      ) as _i4.Future<_i2.Either<_i5.Failure, _i7.ProductModel>>);

  @override
  _i4.Future<_i2.Either<_i5.Failure, void>> deleteProduct(String? id) =>
      (super.noSuchMethod(
        Invocation.method(
          #deleteProduct,
          [id],
        ),
        returnValue: _i4.Future<_i2.Either<_i5.Failure, void>>.value(
            _FakeEither_0<_i5.Failure, void>(
          this,
          Invocation.method(
            #deleteProduct,
            [id],
          ),
        )),
      ) as _i4.Future<_i2.Either<_i5.Failure, void>>);
}
