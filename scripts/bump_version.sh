#!/bin/bash
set -e

# ใช้: ./bump_version.sh [major|minor|patch]
# ถ้าไม่ระบุ จะเพิ่ม patch (v1.2.3 -> v1.2.4)

PART=${1:-patch}

# ดึง tag ล่าสุด
latest_tag=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
echo "🔖 Latest tag: $latest_tag"

# ดึง commit id ที่ tag ล่าสุดชี้อยู่
tag_commit=$(git rev-list -n 1 "$latest_tag" 2>/dev/null || echo "")
# ดึง commit id ล่าสุดใน branch ปัจจุบัน
current_commit=$(git rev-parse HEAD)

# ถ้า commit ล่าสุดมี tag แล้ว ให้หยุดเลย
if [[ "$tag_commit" == "$current_commit" ]]; then
  echo "✅ Current commit already tagged with $latest_tag"
  exit 0
fi

# แยก version ออกมา
version=${latest_tag#v}
IFS='.' read -r major minor patch <<< "$version"

# เพิ่ม version ตาม argument
case $PART in
  major)
    major=$((major + 1)); minor=0; patch=0 ;;
  minor)
    minor=$((minor + 1)); patch=0 ;;
  patch)
    patch=$((patch + 1)) ;;
  *)
    echo "❌ Invalid part: $PART (use major|minor|patch)"
    exit 1 ;;
esac

new_tag="v${major}.${minor}.${patch}"
echo "🚀 New version: $new_tag"

# สร้าง tag และ push
read -p "Push tag $new_tag to remote? [y/N] " confirm
if [[ $confirm =~ ^[Yy]$ ]]; then
  git tag "$new_tag"
  git push origin "$new_tag"
  echo "✅ Tag pushed: $new_tag"
else
  echo "❎ Cancelled"
fi