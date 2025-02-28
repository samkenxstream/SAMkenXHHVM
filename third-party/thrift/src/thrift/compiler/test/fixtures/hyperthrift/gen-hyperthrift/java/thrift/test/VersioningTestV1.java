/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package thrift.test;

import java.util.List;
import java.util.Map;
import java.util.Set;
import javax.annotation.concurrent.Immutable;
import javax.annotation.Nullable;
import com.facebook.hyperthrift.HyperThriftBase;
import com.facebook.hyperthrift.reflect.HyperThriftType;

@Immutable
@HyperThriftType
public class VersioningTestV1 extends HyperThriftBase {
  public static final String TYPE_NAME = "thrift.test.VersioningTestV1";


  @Nullable
  public int begin_in_both() {
    return getFieldValue(0);
  }

  @Nullable
  public String old_string() {
    return getFieldValue(1);
  }

  @Nullable
  public int end_in_both() {
    return getFieldValue(2);
  }



  public static class Builder extends HyperThriftBase.Builder {
    public Builder() {
      super(3);
    }

    public Builder(VersioningTestV1 other) {
      super(other);
    }

    @Nullable
    public int begin_in_both() {
      return getFieldValue(0);
    }

    public Builder VersioningTestV1( int value) {
      setFieldValue(0, value);
      return this;
    }

    @Nullable
    public String old_string() {
      return getFieldValue(1);
    }

    public Builder VersioningTestV1( String value) {
      setFieldValue(1, value);
      return this;
    }

    @Nullable
    public int end_in_both() {
      return getFieldValue(2);
    }

    public Builder VersioningTestV1( int value) {
      setFieldValue(2, value);
      return this;
    }

    public VersioningTestV1 build() {
      Object[] fields = markBuilt();
      VersioningTestV1 instance = new VersioningTestV1();
      instance.init(TYPE_NAME, fields);
      return instance;
    }
  }
}
